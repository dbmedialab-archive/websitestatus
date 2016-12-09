<template>
    <ul class="siteList">
        <li v-for="site in sites" class="site">
            <site v-bind:site="site"></site>
        </li>
    </ul>
</template>

<script>
    export default {
        ready: function () {
            fetch("status/all", {method: "GET"}).then(function(resp) {
                return resp.json();
            }).then((json) => {
                this.sites = json;
            }).catch(function(e) {
                document.body.innerHTML = "Error getting site info<br>" + e;
            });

            let client = new EventSource('http://localhost:8080/events'),
                that = this;

            client.onmessage = function(msg) {
                that.sites = JSON.parse(msg.data);
            }
        },
        data: function () {
            return {
                sites: []
            }
        },
        components: {
            'site': require('./site.vue')
        }
    }
</script>
