---
keywords:
  - grafana
  - schema
title: BarChartPanelCfg kind
---
> Both documentation generation and kinds schemas are in active development and subject to change without prior notice.

## BarChartPanelCfg

#### Maturity: experimental
#### Version: 0.0



| Property           | Type                        | Required | Description |
|--------------------|-----------------------------|----------|-------------|
| `PanelFieldConfig` | [object](#panelfieldconfig) | **Yes**  |             |
| `PanelOptions`     | [object](#paneloptions)     | **Yes**  |             |

### PanelFieldConfig

It extends [AxisConfig](#axisconfig) and [HideableFieldConfig](#hideablefieldconfig).

| Property            | Type                                                      | Required | Description                                                                                                                             |
|---------------------|-----------------------------------------------------------|----------|-----------------------------------------------------------------------------------------------------------------------------------------|
| `axisCenteredZero`  | boolean                                                   | No       | *(Inherited from [AxisConfig](#axisconfig))*                                                                                            |
| `axisColorMode`     | string                                                    | No       | *(Inherited from [AxisConfig](#axisconfig))*<br/>TODO docs<br/>Possible values are: `text`, `series`.                                   |
| `axisGridShow`      | boolean                                                   | No       | *(Inherited from [AxisConfig](#axisconfig))*                                                                                            |
| `axisLabel`         | string                                                    | No       | *(Inherited from [AxisConfig](#axisconfig))*                                                                                            |
| `axisPlacement`     | string                                                    | No       | *(Inherited from [AxisConfig](#axisconfig))*<br/>TODO docs<br/>Possible values are: `auto`, `top`, `right`, `bottom`, `left`, `hidden`. |
| `axisSoftMax`       | number                                                    | No       | *(Inherited from [AxisConfig](#axisconfig))*                                                                                            |
| `axisSoftMin`       | number                                                    | No       | *(Inherited from [AxisConfig](#axisconfig))*                                                                                            |
| `axisWidth`         | number                                                    | No       | *(Inherited from [AxisConfig](#axisconfig))*                                                                                            |
| `fillOpacity`       | integer                                                   | No       | Controls the fill opacity of the bars. Default: `80`.                                                                                   |
| `gradientMode`      | string                                                    | No       | TODO docs<br/>Possible values are: `none`, `opacity`, `hue`, `scheme`.                                                                  |
| `hideFrom`          | [HideSeriesConfig](#hideseriesconfig)                     | No       | *(Inherited from [HideableFieldConfig](#hideablefieldconfig))*<br/>TODO docs                                                            |
| `lineWidth`         | integer                                                   | No       | Controls line width of the bars. Default: `1`.                                                                                          |
| `scaleDistribution` | [ScaleDistributionConfig](#scaledistributionconfig)       | No       | *(Inherited from [AxisConfig](#axisconfig))*<br/>TODO docs                                                                              |
| `thresholdsStyle`   | [GraphThresholdsStyleConfig](#graphthresholdsstyleconfig) | No       | TODO docs                                                                                                                               |

### AxisConfig

TODO docs

| Property            | Type                                                | Required | Description                                                                            |
|---------------------|-----------------------------------------------------|----------|----------------------------------------------------------------------------------------|
| `axisCenteredZero`  | boolean                                             | No       |                                                                                        |
| `axisColorMode`     | string                                              | No       | TODO docs<br/>Possible values are: `text`, `series`.                                   |
| `axisGridShow`      | boolean                                             | No       |                                                                                        |
| `axisLabel`         | string                                              | No       |                                                                                        |
| `axisPlacement`     | string                                              | No       | TODO docs<br/>Possible values are: `auto`, `top`, `right`, `bottom`, `left`, `hidden`. |
| `axisSoftMax`       | number                                              | No       |                                                                                        |
| `axisSoftMin`       | number                                              | No       |                                                                                        |
| `axisWidth`         | number                                              | No       |                                                                                        |
| `scaleDistribution` | [ScaleDistributionConfig](#scaledistributionconfig) | No       | TODO docs                                                                              |

### ScaleDistributionConfig

TODO docs

| Property          | Type   | Required | Description                                                              |
|-------------------|--------|----------|--------------------------------------------------------------------------|
| `type`            | string | **Yes**  | TODO docs<br/>Possible values are: `linear`, `log`, `ordinal`, `symlog`. |
| `linearThreshold` | number | No       |                                                                          |
| `log`             | number | No       |                                                                          |

### GraphThresholdsStyleConfig

TODO docs

| Property | Type   | Required | Description                                                                                               |
|----------|--------|----------|-----------------------------------------------------------------------------------------------------------|
| `mode`   | string | **Yes**  | TODO docs<br/>Possible values are: `off`, `line`, `dashed`, `area`, `line+area`, `dashed+area`, `series`. |

### HideSeriesConfig

TODO docs

| Property  | Type    | Required | Description |
|-----------|---------|----------|-------------|
| `legend`  | boolean | **Yes**  |             |
| `tooltip` | boolean | **Yes**  |             |
| `viz`     | boolean | **Yes**  |             |

### HideableFieldConfig

TODO docs

| Property   | Type                                  | Required | Description |
|------------|---------------------------------------|----------|-------------|
| `hideFrom` | [HideSeriesConfig](#hideseriesconfig) | No       | TODO docs   |

### HideSeriesConfig

TODO docs

| Property  | Type    | Required | Description |
|-----------|---------|----------|-------------|
| `legend`  | boolean | **Yes**  |             |
| `tooltip` | boolean | **Yes**  |             |
| `viz`     | boolean | **Yes**  |             |

### ScaleDistributionConfig

TODO docs

| Property          | Type   | Required | Description                                                              |
|-------------------|--------|----------|--------------------------------------------------------------------------|
| `type`            | string | **Yes**  | TODO docs<br/>Possible values are: `linear`, `log`, `ordinal`, `symlog`. |
| `linearThreshold` | number | No       |                                                                          |
| `log`             | number | No       |                                                                          |

### PanelOptions

It extends [OptionsWithLegend](#optionswithlegend) and [OptionsWithTooltip](#optionswithtooltip) and [OptionsWithTextFormatting](#optionswithtextformatting).

| Property              | Type                                            | Required | Description                                                                                                                        |
|-----------------------|-------------------------------------------------|----------|------------------------------------------------------------------------------------------------------------------------------------|
| `barWidth`            | number                                          | **Yes**  | Controls the width of bars. 1 = Max width, 0 = Min width. Default: `0.97`.                                                         |
| `fullHighlight`       | boolean                                         | **Yes**  | Enables mode which highlights the entire bar area and shows tooltip when cursor<br/>hovers over highlighted area Default: `false`. |
| `groupWidth`          | number                                          | **Yes**  | Controls the width of groups. 1 = max with, 0 = min width. Default: `0.7`.                                                         |
| `legend`              | [VizLegendOptions](#vizlegendoptions)           | **Yes**  | *(Inherited from [OptionsWithLegend](#optionswithlegend))*<br/>TODO docs                                                           |
| `orientation`         | string                                          | **Yes**  | TODO docs<br/>Possible values are: `auto`, `vertical`, `horizontal`.                                                               |
| `showValue`           | string                                          | **Yes**  | TODO docs<br/>Possible values are: `auto`, `never`, `always`.                                                                      |
| `stacking`            | string                                          | **Yes**  | TODO docs<br/>Possible values are: `none`, `normal`, `percent`.                                                                    |
| `tooltip`             | [VizTooltipOptions](#viztooltipoptions)         | **Yes**  | *(Inherited from [OptionsWithTooltip](#optionswithtooltip))*<br/>TODO docs                                                         |
| `xTickLabelMaxLength` | integer                                         | **Yes**  | Sets the max length that a label can have before it is truncated.                                                                  |
| `xTickLabelRotation`  | integer                                         | **Yes**  | Controls the rotation of the x axis labels. Default: `0`.                                                                          |
| `barRadius`           | number                                          | No       | Controls the radius of each bar. Default: `0`.                                                                                     |
| `colorByField`        | string                                          | No       | Use the color value for a sibling field to color each bar value.                                                                   |
| `text`                | [VizTextDisplayOptions](#viztextdisplayoptions) | No       | *(Inherited from [OptionsWithTextFormatting](#optionswithtextformatting))*<br/>TODO docs                                           |
| `xField`              | string                                          | No       | Manually select which field from the dataset to represent the x field.                                                             |
| `xTickLabelSpacing`   | integer                                         | No       | Controls the spacing between x axis labels.<br/>negative values indicate backwards skipping behavior Default: `0`.                 |

### OptionsWithLegend

TODO docs

| Property | Type                                  | Required | Description |
|----------|---------------------------------------|----------|-------------|
| `legend` | [VizLegendOptions](#vizlegendoptions) | **Yes**  | TODO docs   |

### VizLegendOptions

TODO docs

| Property      | Type     | Required | Description                                                                                                                             |
|---------------|----------|----------|-----------------------------------------------------------------------------------------------------------------------------------------|
| `calcs`       | string[] | **Yes**  |                                                                                                                                         |
| `displayMode` | string   | **Yes**  | TODO docs<br/>Note: "hidden" needs to remain as an option for plugins compatibility<br/>Possible values are: `list`, `table`, `hidden`. |
| `placement`   | string   | **Yes**  | TODO docs<br/>Possible values are: `bottom`, `right`.                                                                                   |
| `showLegend`  | boolean  | **Yes**  |                                                                                                                                         |
| `asTable`     | boolean  | No       |                                                                                                                                         |
| `isVisible`   | boolean  | No       |                                                                                                                                         |
| `sortBy`      | string   | No       |                                                                                                                                         |
| `sortDesc`    | boolean  | No       |                                                                                                                                         |
| `width`       | number   | No       |                                                                                                                                         |

### OptionsWithTextFormatting

TODO docs

| Property | Type                                            | Required | Description |
|----------|-------------------------------------------------|----------|-------------|
| `text`   | [VizTextDisplayOptions](#viztextdisplayoptions) | No       | TODO docs   |

### VizTextDisplayOptions

TODO docs

| Property    | Type   | Required | Description              |
|-------------|--------|----------|--------------------------|
| `titleSize` | number | No       | Explicit title text size |
| `valueSize` | number | No       | Explicit value text size |

### OptionsWithTooltip

TODO docs

| Property  | Type                                    | Required | Description |
|-----------|-----------------------------------------|----------|-------------|
| `tooltip` | [VizTooltipOptions](#viztooltipoptions) | **Yes**  | TODO docs   |

### VizTooltipOptions

TODO docs

| Property | Type   | Required | Description                                                   |
|----------|--------|----------|---------------------------------------------------------------|
| `mode`   | string | **Yes**  | TODO docs<br/>Possible values are: `single`, `multi`, `none`. |
| `sort`   | string | **Yes**  | TODO docs<br/>Possible values are: `asc`, `desc`, `none`.     |

### VizLegendOptions

TODO docs

| Property      | Type     | Required | Description                                                                                                                             |
|---------------|----------|----------|-----------------------------------------------------------------------------------------------------------------------------------------|
| `calcs`       | string[] | **Yes**  |                                                                                                                                         |
| `displayMode` | string   | **Yes**  | TODO docs<br/>Note: "hidden" needs to remain as an option for plugins compatibility<br/>Possible values are: `list`, `table`, `hidden`. |
| `placement`   | string   | **Yes**  | TODO docs<br/>Possible values are: `bottom`, `right`.                                                                                   |
| `showLegend`  | boolean  | **Yes**  |                                                                                                                                         |
| `asTable`     | boolean  | No       |                                                                                                                                         |
| `isVisible`   | boolean  | No       |                                                                                                                                         |
| `sortBy`      | string   | No       |                                                                                                                                         |
| `sortDesc`    | boolean  | No       |                                                                                                                                         |
| `width`       | number   | No       |                                                                                                                                         |

### VizTextDisplayOptions

TODO docs

| Property    | Type   | Required | Description              |
|-------------|--------|----------|--------------------------|
| `titleSize` | number | No       | Explicit title text size |
| `valueSize` | number | No       | Explicit value text size |

### VizTooltipOptions

TODO docs

| Property | Type   | Required | Description                                                   |
|----------|--------|----------|---------------------------------------------------------------|
| `mode`   | string | **Yes**  | TODO docs<br/>Possible values are: `single`, `multi`, `none`. |
| `sort`   | string | **Yes**  | TODO docs<br/>Possible values are: `asc`, `desc`, `none`.     |


