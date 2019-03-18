import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';
import './custom.css';
import HelloWorld from './components/HelloWorld/HelloWorld';
import Header from './components/Header/Header';
import { TestHook } from './components/TestHook/TestHook';

class App extends Component {
    render() {
        return ( 
            <div className = "App">
                <header className = "App-header">
                    <img src = {logo} className="App-logo" alt="logo" />
                    <Header/>
                    <HelloWorld user="hé hé"/>

                    <TestHook></TestHook>
                </header>

            </div>
        );
    }
}

export default App;