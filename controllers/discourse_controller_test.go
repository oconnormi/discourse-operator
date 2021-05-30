/*
Copyright 2021 oconnormi.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"time"

	. "github.com/onsi/ginkgo"
)

var _ = Describe("Discourse Controller", func() {
	const (
		DiscourseName      = "test-discourse"
		DiscourseNamespace = "default"

		timeout  = time.Second * 10
		duration = time.Second * 10
		interval = time.Millisecond * 250
	)

	// 	Context("When Creating a Discourse Instance", func() {
	// 		It("Should provide a set of default admin credentials", func() {
	// 			By("Creating a new secret when one doesn't exist", func() {
	// 				ctx := context.Background()
	// 				discourse := &v1alpha1.Discourse{
	// 					TypeMeta: metav1.TypeMeta{
	// 						APIVersion: "discourse.oconnormi.io/v1alpha1",
	// 						Kind:       "Discourse",
	// 					},
	// 					ObjectMeta: metav1.ObjectMeta{
	// 						Name:      DiscourseName,
	// 						Namespace: DiscourseNamespace,
	// 					},
	// 					Spec: v1alpha1.DiscourseSpec{},
	// 				}
	// 				Expect(k8sClient.Create(ctx, discourse)).Should(Succeed())

	// 				// adminSecretLookupKey := types.NamespacedName{Name: DiscourseName + "-admin", Namespace: discourse.Namespace}
	// 				// adminSecret := &v1.Secret{}

	// 				// Eventually(func() bool {
	// 				// 	err := k8sClient.Get(ctx, adminSecretLookupKey, adminSecret)
	// 				// 	if err != nil {
	// 				// 		return false
	// 				// 	}
	// 				// 	return true
	// 				// }, timeout, interval).Should(BeTrue())
	// 				// Expect(adminSecret.Data["admin-user"]).ShouldNot(BeNil())
	// 				// Expect(adminSecret.Data["admin-password"]).ShouldNot(BeNil())
	// 			})
	// 			By("Creating a specific secret when name is provided", func() {
	// 				ctx := context.Background()
	// 				discourse := &v1alpha1.Discourse{
	// 					TypeMeta: metav1.TypeMeta{
	// 						APIVersion: "discourse.oconnormi.io/v1alpha1",
	// 						Kind:       "Discourse",
	// 					},
	// 					ObjectMeta: metav1.ObjectMeta{
	// 						Name:      DiscourseName,
	// 						Namespace: DiscourseNamespace,
	// 					},
	// 					Spec: v1alpha1.DiscourseSpec{
	// 						AdminSecret: "foo-admin",
	// 					},
	// 				}
	// 				Expect(k8sClient.Create(ctx, discourse)).Should(Succeed())

	// 				// adminSecretLookupKey := types.NamespacedName{Name: "foo-admin", Namespace: discourse.Namespace}
	// 				// adminSecret := &v1.Secret{}

	// 				// Eventually(func() bool {
	// 				// 	err := k8sClient.Get(ctx, adminSecretLookupKey, adminSecret)
	// 				// 	if err != nil {
	// 				// 		return false
	// 				// 	}
	// 				// 	return true
	// 				// }, timeout, interval).Should(BeTrue())
	// 				// Expect(adminSecret.Data["admin-user"]).ShouldNot(BeNil())
	// 				// Expect(adminSecret.Data["admin-password"]).ShouldNot(BeNil())
	// 			})
	// 		})
	// 	})
	// 	It("Should initialize discourse", func() {
	// 		By("Creating a deployment using the default image", func() {
	// 			ctx := context.Background()
	// 			discourse := &v1alpha1.Discourse{
	// 				TypeMeta: metav1.TypeMeta{
	// 					APIVersion: "discourse.oconnormi.io/v1alpha1",
	// 					Kind:       "Discourse",
	// 				},
	// 				ObjectMeta: metav1.ObjectMeta{
	// 					Name:      DiscourseName,
	// 					Namespace: DiscourseNamespace,
	// 				},
	// 				Spec: v1alpha1.DiscourseSpec{},
	// 			}
	// 			Expect(k8sClient.Create(ctx, discourse)).Should(Succeed())

	// 			deploymentLookupKey := types.NamespacedName{Name: DiscourseName, Namespace: DiscourseNamespace}
	// 			deployment := appsv1.Deployment{}

	// 			Eventually(func() bool {
	// 				err := k8sClient.Get(ctx, deploymentLookupKey, &deployment)
	// 				if err != nil {
	// 					return false
	// 				}
	// 				return true
	// 			}, timeout, interval).Should(BeTrue())
	// 			// Expect(deployment.Spec.Template.Spec.Containers).Should(ContainSubstring("discourse:latest"))
	// 		})
	// 		By("Creating a deployment using the specified image", func() {
	// 			ctx := context.Background()
	// 			discourse := &v1alpha1.Discourse{
	// 				TypeMeta: metav1.TypeMeta{
	// 					APIVersion: "discourse.oconnormi.io/v1alpha1",
	// 					Kind:       "Discourse",
	// 				},
	// 				ObjectMeta: metav1.ObjectMeta{
	// 					Name:      DiscourseName,
	// 					Namespace: DiscourseNamespace,
	// 				},
	// 				Spec: v1alpha1.DiscourseSpec{
	// 					Image: "foo:bar",
	// 				},
	// 			}
	// 			Expect(k8sClient.Create(ctx, discourse)).Should(Succeed())

	// 			deploymentLookupKey := types.NamespacedName{Name: DiscourseName, Namespace: DiscourseNamespace}
	// 			deployment := appsv1.Deployment{}

	// 			Eventually(func() bool {
	// 				err := k8sClient.Get(ctx, deploymentLookupKey, &deployment)
	// 				if err != nil {
	// 					return false
	// 				}
	// 				return true
	// 			}, timeout, interval).Should(BeTrue())
	// 			// Expect(deployment.Spec.Template.Spec.Containers).Should(ContainSubstring("foo:bar"))
	// 		})
	// 		By("Installing specified plugins", func() {
	// 			ctx := context.Background()
	// 			discourse := &v1alpha1.Discourse{
	// 				TypeMeta: metav1.TypeMeta{
	// 					APIVersion: "discourse.oconnormi.io/v1alpha1",
	// 					Kind:       "Discourse",
	// 				},
	// 				ObjectMeta: metav1.ObjectMeta{
	// 					Name:      DiscourseName,
	// 					Namespace: DiscourseNamespace,
	// 				},
	// 				Spec: v1alpha1.DiscourseSpec{
	// 					Plugins: v1alpha1.Plugins{
	// 						Installed: []string{"foo", "bar", "baz"},
	// 					},
	// 				},
	// 			}
	// 			Expect(k8sClient.Create(ctx, discourse)).Should(Succeed())

	// 			deploymentLookupKey := types.NamespacedName{Name: DiscourseName, Namespace: DiscourseNamespace}
	// 			deployment := appsv1.Deployment{}

	// 			Eventually(func() bool {
	// 				err := k8sClient.Get(ctx, deploymentLookupKey, &deployment)
	// 				if err != nil {
	// 					return false
	// 				}
	// 				return true
	// 			}, timeout, interval).Should(BeTrue())
	// 			// Expect(deployment.Spec.Template.Spec.Containers).Should(ContainSubstring("plugin-manager"))
	// 		})
	// 	})
})
