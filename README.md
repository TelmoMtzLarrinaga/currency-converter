# Currency Converter

[![Super-Linter](https://github.com/TelmoMtzLarrinaga/currency-converter/actions/workflows/linter.yaml/badge.svg)](https://github.com/marketplace/actions/super-linter)

We'll need to convert between a number of base currencies.
The user will be able to select the different currencies through the `huh?` package and we'll need to use a third party API in order to obtain currency conversion data.
Here are some suggested packages you may want to use.
`net/http` for http request to the currency exchange API.
`encoding/json` in order to marshal the data for the API.
`charmbracelet/huh` for the TUI interface form.

## Goals

:white_large_square: Use a thid party API to obtain currency conversion data.

:white_large_square: Use the `huh?` package to create a TUI.

:white_large_square: Convert between a list of base currencies.

:white_large_square: Provide a 1.0.0 release through GitHub.

:white_large_square: Make useful workflows to provide this commodity.

>[!NOTE]
> Please refer to the following GitHub [repository](https://github.com/dreamsofcode-io/goprojects/tree/main), which is where
> this project was proposed.
