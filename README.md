# Happy&Fy - redirect-service

## Configuration

### ENVs

* REDIRECT_CONFIG
  * Json-Configuration as String
* REDIRECT_CONFIG_FILE
  * Path to Json-Configuration File

## Json-Configuration

Example:
```json
{
  "rules": [
    {
      "prio": 0,
      "type": "accept-language",
      "value": "de",
      "target": "https://de.example.com/{path}"
    }
  ],
  "fallback": {
    "type": "fallback",
    "target": "https://www.example.com/{path}"
  }
}
```

### Rules contain:
* `prio` the priority of the rule, the higher, the earlier
* `type` the type of the rule, currently supported:
  * `accept-language` - If value match the Accepted-Header header, the redirect
  * If you have a wish, simply open an issue
* `value` a value depends on type
* `target` target of the redirect

### Fallback:
Fallback will be executed at very last, if no rule matched. `type` must be `fallback`.

### Placeholder
* `{path}` used path of the base url of the redirect-service
* If you have a wish, simply open an issue