import React, { Component } from 'react';

export default class HelloWorld extends Component {
        constructor(props) {
                super(props);
                this.state = {
                     online: true,
                     counter: 0
                }
        }

        callbackclick() {
                console.log("erreur ...");
        }

        render() {
                const user = this.props.user;
                return (
                        <div>
                                <p> Hello {user}, you are {!this.state.online || 'online'} ! </p>
                                <p>{this.state.counter}</p>
                                <button onClick={
                                        () => {
                                                this.setState({counter: this.state.counter + 1})                                                
                                        }
                                }>increment counter</button>
                                
                        </div>
                        
                )
      }
}
