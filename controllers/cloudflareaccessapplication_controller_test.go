//go:build integration

package controllers

import (
	"context"
	"time"

	v1alpha1 "github.com/bojanzelic/cloudflare-zero-trust-operator/api/v1alpha1"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var _ = Describe("CloudflareAccessApplication controller", Ordered, func() {
	BeforeAll(func() {
		ctx := context.Background()

		By("Removing all existing access apps")
		apps, err := api.AccessApplications(ctx)
		Expect(err).To(Not(HaveOccurred()))
		for _, app := range apps {
			err = api.DeleteAccessApplication(ctx, app.ID)
			Expect(err).To(Not(HaveOccurred()))
		}
	})

	Context("CloudflareAccessApplication controller test", func() {

		const cloudflareName = "cloudflare-app"

		ctx := context.Background()

		namespace := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name:      cloudflareName,
				Namespace: cloudflareName,
			},
		}

		typeNamespaceName := types.NamespacedName{Name: cloudflareName, Namespace: cloudflareName}

		BeforeEach(func() {
			By("Creating the Namespace to perform the tests")
			k8sClient.Create(ctx, namespace)
			// ignore error because of https://book.kubebuilder.io/reference/envtest.html#namespace-usage-limitation
			//Expect(err).To(Not(HaveOccurred()))
		})

		// AfterEach(func() {
		// 	By("Deleting the Namespace to perform the tests")
		// 	//_ = k8sClient.Delete(ctx, namespace)
		// })

		It("should successfully reconcile CloudflareAccessApplication policies", func() {
			By("Creating the custom resource for the Kind CloudflareAccessApplication")
			typeNamespaceName := types.NamespacedName{Name: "cloudflare-app-two", Namespace: cloudflareName}
			apps := &v1alpha1.CloudflareAccessApplication{
				ObjectMeta: metav1.ObjectMeta{
					Name:      typeNamespaceName.Name,
					Namespace: namespace.Name,
				},
				Spec: v1alpha1.CloudflareAccessApplicationSpec{
					Name:   "integration policies test",
					Domain: "integration-policies.cf-operator-tests.uk",
					Policies: v1alpha1.CloudflareAccessPolicyList{
						{
							Name:     "integration_test",
							Decision: "allow",
							Include: []v1alpha1.CloudFlareAccessGroupRule{
								{
									Emails: []string{"testemail@cf-operator-tests.uk", "testemail2@cf-operator-tests.uk"},
								},
								{
									EmailDomains: []string{"cf-operator-tests.uk"},
								},
							},
						},
					},
				},
			}
			err := k8sClient.Create(ctx, apps)
			Expect(err).To(Not(HaveOccurred()))

			By("Reconciling the custom resource created")
			accessGroupReconciler := &CloudflareAccessApplicationReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err = accessGroupReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespaceName,
			})
			Expect(err).To(Not(HaveOccurred()))

			found := &v1alpha1.CloudflareAccessApplication{}
			By("Checking the latest Status should have the ID of the resource")
			Eventually(func() string {
				found = &v1alpha1.CloudflareAccessApplication{}
				k8sClient.Get(ctx, typeNamespaceName, found)
				return found.Status.AccessApplicationID
			}, time.Minute, time.Second).Should(Not(BeEmpty()))

			By("Cloudflare resource should equal the spec")
			cfResource, err := api.AccessPolicies(ctx, found.Status.AccessApplicationID)
			Expect(err).To(Not(HaveOccurred()))
			Expect(cfResource[0].Name).To(Equal(found.Spec.Policies[0].Name))
			Expect(cfResource[0].Include[0].(map[string]interface{})["email"].(map[string]interface{})["email"]).To(Equal(found.Spec.Policies[0].Include[0].Emails[0]))
			Expect(cfResource[0].Include[1].(map[string]interface{})["email"].(map[string]interface{})["email"]).To(Equal(found.Spec.Policies[0].Include[0].Emails[1]))
			Expect(cfResource[0].Include[2].(map[string]interface{})["email_domain"].(map[string]interface{})["domain"]).To(Equal(found.Spec.Policies[0].Include[1].EmailDomains[0]))

			By("changing a policy")
			found.Spec.Policies[0].Name = "updated_policy"
			found.Spec.Policies[0].Include[0].Emails[0] = "testemail3@cf-operator-tests.uk"
			err = k8sClient.Update(ctx, found)
			Expect(err).To(Not(HaveOccurred()))

			By("Reconciling the updated custom resource")
			_, err = accessGroupReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespaceName,
			})
			Expect(err).To(Not(HaveOccurred()))

			By("Cloudflare resource should equal the spec")
			cfResource, err = api.AccessPolicies(ctx, found.Status.AccessApplicationID)
			Expect(err).To(Not(HaveOccurred()))
			Expect(cfResource[0].Name).To(Equal(found.Spec.Policies[0].Name))
			Expect(cfResource[0].Include[0].(map[string]interface{})["email"].(map[string]interface{})["email"]).To(Equal(found.Spec.Policies[0].Include[0].Emails[0]))
			Expect(cfResource[0].Include[1].(map[string]interface{})["email"].(map[string]interface{})["email"]).To(Equal(found.Spec.Policies[0].Include[0].Emails[1]))
			Expect(cfResource[0].Include[2].(map[string]interface{})["email_domain"].(map[string]interface{})["domain"]).To(Equal(found.Spec.Policies[0].Include[1].EmailDomains[0]))

		})

		It("should successfully reconcile a custom resource for CloudflareAccessApplication", func() {
			By("Creating the custom resource for the Kind CloudflareAccessApplication")
			apps := &v1alpha1.CloudflareAccessApplication{
				ObjectMeta: metav1.ObjectMeta{
					Name:      cloudflareName,
					Namespace: namespace.Name,
				},
				Spec: v1alpha1.CloudflareAccessApplicationSpec{
					Name:   "integration test",
					Domain: "integration.cf-operator-tests.uk",
				},
			}
			err := k8sClient.Create(ctx, apps)
			Expect(err).To(Not(HaveOccurred()))

			By("Checking if the custom resource was successfully created")
			Eventually(func() error {
				found := &v1alpha1.CloudflareAccessApplication{}
				return k8sClient.Get(ctx, typeNamespaceName, found)
			}, time.Minute, time.Second).Should(Succeed())

			By("Reconciling the custom resource created")
			accessGroupReconciler := &CloudflareAccessApplicationReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err = accessGroupReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespaceName,
			})
			Expect(err).To(Not(HaveOccurred()))

			found := &v1alpha1.CloudflareAccessApplication{}
			By("Checking the latest Status should have the ID of the resource")
			Eventually(func() string {
				found = &v1alpha1.CloudflareAccessApplication{}
				k8sClient.Get(ctx, typeNamespaceName, found)
				return found.Status.AccessApplicationID
			}, time.Minute, time.Second).Should(Not(BeEmpty()))

			By("Cloudflare resource should equal the spec")
			cfResource, err := api.AccessApplication(ctx, found.Status.AccessApplicationID)
			Expect(err).To(Not(HaveOccurred()))
			Expect(cfResource.Name).To(Equal(found.Spec.Name))

			By("Updating the name of the resource")
			found.Spec.Name = "updated name"
			k8sClient.Update(ctx, found)
			Expect(err).To(Not(HaveOccurred()))

			By("Reconciling the updated resource")
			_, err = accessGroupReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespaceName,
			})
			Expect(err).To(Not(HaveOccurred()))

			By("Cloudflare resource should equal the updated spec")
			cfResource, err = api.AccessApplication(ctx, found.Status.AccessApplicationID)
			Expect(err).To(Not(HaveOccurred()))
			Expect(cfResource.Name).To(Equal(found.Spec.Name))
		})
	})
})