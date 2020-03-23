import Vue from 'vue'
import VueRouter from 'vue-router'
import Vuelidate from 'vuelidate'
import App from './App.vue'
import { routes } from "./routes"
import { store } from './store/store'

import Buefy from 'buefy'
Vue.use(Buefy)

Vue.use(Vuelidate)
Vue.use(VueRouter)

window.axios = require('axios')

const router = new VueRouter({
    routes
})

new Vue({
    el: '#app',
    router,
    store,
    render: h => h(App)
})