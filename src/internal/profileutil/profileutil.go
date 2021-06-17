// Profileutil contains functionality to export performance information to external systems.
package profileutil

import (
	"cloud.google.com/go/profiler"
	"github.com/pachyderm/pachyderm/v2/src/version"
)

// StartCloudProfiler enables Google Cloud Profiler and begins exporting profiles.  Configuration
// information is read from the instance's GCE metadata server, so this requires running on GCP.
// Docs: https://cloud.google.com/profiler/docs
//
// Service is the name of this binary (pachd, worker, etc.).
func StartCloudProfiler(service string) error {
	return profiler.Start(profiler.Config{
		Service:        service,
		ServiceVersion: version.PrettyPrintVersion(version.Version),
		MutexProfiling: true,
		DebugLogging:   true, // XXX: disable before merge.
	})
}
