<template>
    <span>{{ site.Site.Name }} </span>
    <span class="{{ getStatusClass() }}">Status: {{ site.Status }} </span>
    <span>Responsetime: {{ getResponseTime() }}ms </span>
    <span>Updated: {{ site.Updated }} </span>
    <span>Size: {{ site.Size }}KB </span>
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
            }
        }
    }
</script>
