<template>
    <h1 class="title">{{ site.Site.Name }}</h1>
    <p>Status: <span class="label {{ getStatusClass() }}">{{ site.Status }}</span> </p>
    <p>Responsetime: <span class="label {{ getResponseTimeClass() }}">{{ getResponseTime() }}ms</span> </p>
    <p>Size: <span>{{ site.Size }}KB</span> </p>
    <p>Updated: <span>{{ site.Updated }}</span> </p>
</template>

<script>
    export default {
        props: ['site'],
        methods: {
            getResponseTime: function () {
                return Math.round(this.site.ResponseTime * 100) / 100
            },
            getStatusClass: function () {
                switch(this.site.Status) {
                    case 200:
                        return 'success'
                    default:
                        return 'error'
                }
            },
            getResponseTimeClass: function () {
                let time = this.getResponseTime()
                if (time < 500) {
                    return 'success'
                } else if (time < 2000) {
                    return 'warning'
                }
                return 'error'
            }
        }
    }
</script>
