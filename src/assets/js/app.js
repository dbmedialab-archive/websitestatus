'use strict';

var Vue = require('Vue');

const app = new Vue({
    el: '#main',
    components: {
        'sitelist': require('./components/sitelist.vue')
    }
});
