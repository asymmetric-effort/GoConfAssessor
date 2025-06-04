goConfAssessor
==============

## Assess a set of files using a manifest and write reports
```shell
assess \
  --manifest /path/to/rootAssessment.yaml \
  --source /path/to/configs/ \
  --report-dir /path/to/output-reports/
```

## Create a skeleton manifest YAML file
```shell
create-manifest --name filename.yaml
```

## Analyze generated reports to produce a score
```shell
analyze --report-dir /path/to/output-reports/ --output score.yaml
```

## Verify manifest
```shell
manifest-verify --manifest /path/to/rootAssessment.yaml 
```
