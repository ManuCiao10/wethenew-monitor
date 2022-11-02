<h2>Open Source Wethenew Monitor Based on GoLang</h2>

_A Go Web Application to srape and monitoring an API_

<h3>The module contains the following Features</h3>

<h4>- Ability to send a HTTPS request to Wethenew API every seconds(the value is decide by the client)</h4>
<h4>- Run the web application on Docker</h4>

<h4>- Ability to send the new items to the client via Webhook message to Discord</h4>
<h4>- Comparing the old items with the new-one to send</h4>


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

