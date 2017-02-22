// deepcopy-gen is a tool for auto-generating DeepCopy functions.
//
// Structs in the input directories with the below line in their comments
// will be ignored during generation.
// // +gencopy=false
package main

import (
	"strings"

	"github.com/golang/glog"

	"k8s.io/gengo/args"
	"k8s.io/gengo/examples/deepcopy-gen/generators"
	"k8s.io/gengo/generator"
)

func main() {
	arguments := args.Default()

	// Override defaults. These are Kubernetes specific input locations.
	arguments.InputDirs = []string{
		"k8s.io/kubernetes/pkg/api",
		"k8s.io/kubernetes/pkg/api/unversioned",
		"k8s.io/kubernetes/pkg/api/v1",
		"k8s.io/kubernetes/pkg/apis/authorization",
		"k8s.io/kubernetes/pkg/apis/authorization/v1beta1",
		"k8s.io/kubernetes/pkg/apis/autoscaling",
		"k8s.io/kubernetes/pkg/apis/autoscaling/v1",
		"k8s.io/kubernetes/pkg/apis/batch",
		"k8s.io/kubernetes/pkg/apis/batch/v1",
		"k8s.io/kubernetes/pkg/apis/componentconfig",
		"k8s.io/kubernetes/pkg/apis/componentconfig/v1alpha1",
		"k8s.io/kubernetes/pkg/apis/extensions",
		"k8s.io/kubernetes/pkg/apis/extensions/v1beta1",
		"k8s.io/kubernetes/pkg/runtime",
		"k8s.io/kubernetes/pkg/util/intstr",
		"k8s.io/kubernetes/pkg/util/sets",
		"github.com/openshift/origin/pkg/authorization/api/v1",
		"github.com/openshift/origin/pkg/authorization/api",
		"github.com/openshift/origin/pkg/build/api/v1",
		"github.com/openshift/origin/pkg/build/api",
		"github.com/openshift/origin/pkg/deploy/api/v1",
		"github.com/openshift/origin/pkg/deploy/api",
		"github.com/openshift/origin/pkg/image/api/v1",
		"github.com/openshift/origin/pkg/image/api",
		"github.com/openshift/origin/pkg/oauth/api/v1",
		"github.com/openshift/origin/pkg/oauth/api",
		"github.com/openshift/origin/pkg/project/api/v1",
		"github.com/openshift/origin/pkg/project/api",
		"github.com/openshift/origin/pkg/quota/api/v1",
		"github.com/openshift/origin/pkg/quota/api",
		"github.com/openshift/origin/pkg/route/api/v1",
		"github.com/openshift/origin/pkg/route/api",
		"github.com/openshift/origin/pkg/sdn/api/v1",
		"github.com/openshift/origin/pkg/sdn/api",
		"github.com/openshift/origin/pkg/template/api/v1",
		"github.com/openshift/origin/pkg/template/api",
		"github.com/openshift/origin/pkg/user/api/v1",
		"github.com/openshift/origin/pkg/user/api",
		"github.com/openshift/origin/pkg/security/api/v1",
		"github.com/openshift/origin/pkg/security/api",
	}

	arguments.GeneratedBuildTag = "ignore_autogenerated_openshift"
	arguments.GoHeaderFilePath = "hack/boilerplate.txt"
	arguments.OutputFileBaseName = "zz_generated.deepcopy"
	arguments.CustomArgs = &generators.CustomArgs{
		BoundingDirs: []string{
			"k8s.io/kubernetes",
			"github.com/openshift/origin",
		},
	}

	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		func(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
			pkgs := generators.Packages(context, arguments)
			var include generator.Packages
			for _, pkg := range pkgs {
				if strings.HasPrefix(pkg.Path(), "k8s.io/") {
					continue
				}
				include = append(include, pkg)
			}
			return include
		},
	); err != nil {
		glog.Fatalf("Error: %v", err)
	}
	glog.Info("Completed successfully.")
}