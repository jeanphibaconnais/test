import Vue from 'vue';
import App2 from './components/App2';
import App3 from './components/App3';
import App4 from './components/App4';

var app = new Vue({
	el: '#app3',
	template: '<app3/>',
	components: { App2, App3, App4 }
});
