import Vue from 'vue';
import App2 from './components/App2';
import App3 from './components/App3';
import App4 from './components/App4';
import Todo from './components/2/Todo';

var app = new Vue({
	el: '#app4',
	template: '<app4/>',
	components: { App2, App3, App4, Todo }
});
