import React, { Component } from 'react';
// import logo from './logo.svg';
import './App.css';
import './custom.css';
import HelloWorld from './components/HelloWorld/HelloWorld';
import Header from './components/Header/Header';
import { TestHook } from './components/TestHook/TestHook';
import TodoList from './components/TodoList/TodoList';
import TodoListComplex from './components/TodoListComplex/TodoListComplex';

// Redux
import { createStore, combineReducers } from 'redux';
import { reducer as formReducer } from 'redux-form';
import reducer from './components/reducer';
import { Provider } from 'react-redux';

const reducers = {
    reducer,
    form: formReducer,
  }
const reduc = combineReducers(reducers);
const store = createStore(reduc);

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
                    </header>

                </div>
            </Provider>
        );
    }
}

export default App;