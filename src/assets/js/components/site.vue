<template>
    <span>{{ Site.Name }} </span>
    <span>Status: {{ Status }} </span>
    <span>Responsetime: {{ ResponseTime }} </span>
    <span>Updated: {{ Updated }} </span>
</template>

<script>
    export default {
        ready: function () {
            console.log(this.data);
        },
        props: function () {
            return [
                'Site'
            ]
        }
    }
</script>
