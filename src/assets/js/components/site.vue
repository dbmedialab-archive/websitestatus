<template>
    <h1 class="title">{{ site.Site.Name }}</h1>
    <table>
        <tr>
            <td>Status:</td>
            <td><span class="label {{ getStatusClass() }}">{{ site.Status }}</span></td>
        </tr>
        <tr>
            <td>Responsetime:</td>
            <td><span class="label {{ getResponseTimeClass() }}">{{ getResponseTime() }}ms</span></td>
        </tr>
        <tr>
            <td>Size:</td>
            <td><span>{{ site.Size }}KB</span></td>
        </tr>
        <tr>
            <td>Updated:</td>
            <td><span>{{ site.Updated }}</span></td>
        </tr>
    </table>
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
