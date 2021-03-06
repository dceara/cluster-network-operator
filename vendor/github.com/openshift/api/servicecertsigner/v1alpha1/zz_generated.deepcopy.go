// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIServiceCABundleInjectorConfig) DeepCopyInto(out *APIServiceCABundleInjectorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.GenericControllerConfig.DeepCopyInto(&out.GenericControllerConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIServiceCABundleInjectorConfig.
func (in *APIServiceCABundleInjectorConfig) DeepCopy() *APIServiceCABundleInjectorConfig {
	if in == nil {
		return nil
	}
	out := new(APIServiceCABundleInjectorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *APIServiceCABundleInjectorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapCABundleInjectorConfig) DeepCopyInto(out *ConfigMapCABundleInjectorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.GenericControllerConfig.DeepCopyInto(&out.GenericControllerConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapCABundleInjectorConfig.
func (in *ConfigMapCABundleInjectorConfig) DeepCopy() *ConfigMapCABundleInjectorConfig {
	if in == nil {
		return nil
	}
	out := new(ConfigMapCABundleInjectorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConfigMapCABundleInjectorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceCertSignerOperatorConfig) DeepCopyInto(out *ServiceCertSignerOperatorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceCertSignerOperatorConfig.
func (in *ServiceCertSignerOperatorConfig) DeepCopy() *ServiceCertSignerOperatorConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceCertSignerOperatorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceCertSignerOperatorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceCertSignerOperatorConfigList) DeepCopyInto(out *ServiceCertSignerOperatorConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceCertSignerOperatorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceCertSignerOperatorConfigList.
func (in *ServiceCertSignerOperatorConfigList) DeepCopy() *ServiceCertSignerOperatorConfigList {
	if in == nil {
		return nil
	}
	out := new(ServiceCertSignerOperatorConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceCertSignerOperatorConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceCertSignerOperatorConfigSpec) DeepCopyInto(out *ServiceCertSignerOperatorConfigSpec) {
	*out = *in
	in.OperatorSpec.DeepCopyInto(&out.OperatorSpec)
	in.ServiceServingCertSignerConfig.DeepCopyInto(&out.ServiceServingCertSignerConfig)
	in.APIServiceCABundleInjectorConfig.DeepCopyInto(&out.APIServiceCABundleInjectorConfig)
	in.ConfigMapCABundleInjectorConfig.DeepCopyInto(&out.ConfigMapCABundleInjectorConfig)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceCertSignerOperatorConfigSpec.
func (in *ServiceCertSignerOperatorConfigSpec) DeepCopy() *ServiceCertSignerOperatorConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceCertSignerOperatorConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceCertSignerOperatorConfigStatus) DeepCopyInto(out *ServiceCertSignerOperatorConfigStatus) {
	*out = *in
	in.OperatorStatus.DeepCopyInto(&out.OperatorStatus)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceCertSignerOperatorConfigStatus.
func (in *ServiceCertSignerOperatorConfigStatus) DeepCopy() *ServiceCertSignerOperatorConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceCertSignerOperatorConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceServingCertSignerConfig) DeepCopyInto(out *ServiceServingCertSignerConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.GenericControllerConfig.DeepCopyInto(&out.GenericControllerConfig)
	out.Signer = in.Signer
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceServingCertSignerConfig.
func (in *ServiceServingCertSignerConfig) DeepCopy() *ServiceServingCertSignerConfig {
	if in == nil {
		return nil
	}
	out := new(ServiceServingCertSignerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceServingCertSignerConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
