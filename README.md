<h2>Open Source Wethenew Monitor Based on GoLang</h2>

_A Go Web Application to scrape and monitoring an API_


<h3>The module contains the following Features</h3>

- Ability to send a [HTTPS](https://pkg.go.dev/net/http) request to Wethenew API every seconds
- Supporting TLS client
- Run the web application on Docker

- Ability to send the new items to the client via Webhook message to Discord
- Comparing the old items with the new-one to send


```go
for idx, v := range new_id.Results {
    if !Contains(Slice, v.ID) {
        Slice = append(Slice, v.ID)
        discord.Webhook(new_id, idx)
        continue
    }
}
```
<h3>How to use</h3>

<h3>Installation</h3>

