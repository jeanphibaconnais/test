import React, { Component } from 'react';
import './App.css';
import './custom.css';
import HelloWorld from './components/HelloWorld/HelloWorld';
import Header from './components/Header/Header';
import { TestHook } from './components/TestHook/TestHook';
import TodoList from './components/TodoList/TodoList';
import TodoListComplex from './components/TodoListComplex/TodoListComplex';

// Redux
import { createStore } from 'redux';
import reducer from './components/CounterRedux/redux/reducer';
import { Provider } from 'react-redux';
import CounterRedux from './components/CounterRedux/CounterRedux';

// function combineReducer allows to have multiple reducers.
const store = createStore(reducer);

class App extends Component {
    render() {
        return ( 
            <Provider store={store}>
                <div className = "App">
                    <header className = "App-header">
                        <Header/>
                        
                        <HelloWorld user="hé hé"/>

                        <TestHook></TestHook>

                        <TodoList></TodoList>

                        <TodoListComplex></TodoListComplex>

                        <CounterRedux></CounterRedux>
                    </header>
                </div>
            </Provider>
        );
    }
}

export default App;