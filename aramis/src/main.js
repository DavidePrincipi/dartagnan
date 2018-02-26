import Vue from 'vue'
import VueResource from 'vue-resource'
import VueI18n from 'vue-i18n'
import VueClipboard from 'vue-clipboards'

import App from './App'
import router from './routes/router'
import languages from './i18n/lang'
import filters from './filters/filters'

Vue.config.productionTip = true
Vue.use(VueResource)
Vue.use(VueI18n)
Vue.use(VueClipboard)

window.$ = window.jQuery = require('patternfly/node_modules/jquery/dist/jquery.min.js')
require('patternfly/node_modules/bootstrap/dist/js/bootstrap.min.js');
require('patternfly/dist/js/patternfly.min');
require('patternfly/dist/css/patternfly.min.css');
require('patternfly/dist/css/patternfly-additions.min.css');

// configure i18n
var langConf = languages.initLang()
const i18n = new VueI18n({
  locale: langConf.locale,
  messages: langConf.messages,
})
var moment = require("patternfly/node_modules/moment/moment.js")
moment.locale(langConf.locale)

// import filters
Vue.filter('byteFormat', filters.byteFormat)
Vue.filter('secondsInReadable', filters.secondsInReadable)
Vue.filter('formatDate', filters.formatDate)
Vue.filter('dateFromNow', filters.dateFromNow)

// init Vue app
new Vue({
  el: '#app',
  router,
  i18n,
  render: (h) => h(App),
  currentLocale: langConf.locale,
  api_host: window.location.host
})
