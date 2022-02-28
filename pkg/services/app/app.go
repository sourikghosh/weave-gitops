package app

import (
	"context"

	"github.com/benbjohnson/clock"
	"github.com/fluxcd/go-git-providers/gitprovider"
	wego "github.com/weaveworks/weave-gitops/api/v1alpha1"
	"github.com/weaveworks/weave-gitops/pkg/gitproviders"
	"github.com/weaveworks/weave-gitops/pkg/kube"
	"github.com/weaveworks/weave-gitops/pkg/logger"
	"k8s.io/apimachinery/pkg/types"
)

// AppService entity that manages applications
type AppService interface {
	// Get returns a given applicaiton
	Get(name types.NamespacedName) (*wego.Application, error)
	// GetCommits returns a list of commits for an application
	GetCommits(gitProvider gitproviders.GitProvider, params CommitParams, application *wego.Application) ([]gitprovider.Commit, error)
	// Sync trigger reconciliation loop for an application
	Sync(params SyncParams) error
}

type AppSvc struct {
	Context context.Context
	Kube    kube.Kube
	Logger  logger.Logger
	Clock   clock.Clock
}

func New(ctx context.Context, logger logger.Logger, kube kube.Kube) AppService {
	return &AppSvc{
		Context: ctx,
		Kube:    kube,
		Logger:  logger,
		Clock:   clock.New(),
	}
}

// Make sure App implements all the required methods.
var _ AppService = &AppSvc{}
