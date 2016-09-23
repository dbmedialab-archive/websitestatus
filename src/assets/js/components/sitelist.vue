<template>
    <ul class="siteList">
        <li v-for="site in sites" class="site">
            <span>{{ site.Site.Name }} </span>
            <span>Status: {{ site.Status }} </span>
            <span>Responsetime: {{ site.ResponseTime }} </span>
            <span>Updated: {{ site.Updated }} </span>
        </li>
    </ul>
</template>

<script>
    export default {
        ready: function () {
            let client = new EventSource('http://localhost:8080/events'),
                that = this;

            client.onmessage = function(msg) {
                that.sites = JSON.parse(msg.data);
                console.log(that.sites);
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
