import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './assets/tailwind.css'

Vue.config.productionTip = false

import ARDropdown from '@/components/Dropdown.vue'
Vue.component('ar-dropdown', ARDropdown)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
