package utils

import (
	"fmt"
	"github.com/sam-caldwell/GoConfAssessor/pkg/logger"
	"github.com/sam-caldwell/GoConfAssessor/pkg/parsers"
	"github.com/sam-caldwell/directory"
	"github.com/sam-caldwell/file"
	"github.com/sirupsen/logrus"
)

func VerifyArgs(log *logrus.Logger, debug *bool, manifestFile, sourceDir, reportDir, confType *string) (err error) {
	if err = parsers.IsValidParser(*confType); err != nil {
		return err
	}
	if !file.Existsp(manifestFile) {
		return fmt.Errorf("file not found (%s)", *manifestFile)
	}
	if err = logger.SetLevel(logger.IsDebug(*debug)); err != nil {
		log.Error(err)
	}
	if !directory.Exists(*sourceDir) {
		log.Fatalf("source directory not found: '%s'", *sourceDir)
	}
	if !directory.Exists(*reportDir) {
		log.Fatalf("report directory not found: '%s'", *reportDir)
	}
	return nil
}
