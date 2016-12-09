<template>
    <h1 class="title">{{ site.site.name }}</h1>
    <table>
        <tr>
            <td>Status:</td>
            <td><span class="label {{ getStatusClass() }}">{{ site.status }}</span></td>
        </tr>
        <tr>
            <td>Responsetime:</td>
            <td><span class="label {{ getResponseTimeClass() }}">{{ getResponseTime() }}ms</span></td>
        </tr>
        <tr>
            <td>Size:</td>
            <td><span>{{ site.size }}KB</span></td>
        </tr>
        <tr>
            <td>Updated:</td>
            <td><span>{{ site.updated }}</span></td>
        </tr>
    </table>
</template>

<script>
    export default {
        props: ['site'],
        data () {
            return {}
        },
        methods: {
            getResponseTime: function () {
                return Math.round(this.site.responsetime * 100) / 100
            },
            getStatusClass: function () {
                switch(this.site.status) {
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
