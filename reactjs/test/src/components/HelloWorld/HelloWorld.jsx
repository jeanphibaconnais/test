import React, { Component } from 'react';

export default class HelloWorld extends Component {
    constructor(props) {
        super(props);
        this.state = {
            online: true,
            counter: 0
        };
    }

    callbackclick() {
        console.log('erreur ...');
    }
    
    render() {
        const user = this.props.user;
        let classNameSpan = 'span';

        return (
            <div>
                <h3>Component with props</h3>
                <p>
                    hello {user}, you are {!this.state.online || 'online'} !{' '}
                </p>

                <h3>Change state</h3>
                <span className={classNameSpan}>Counter : {this.state.counter}</span>  
                <button onClick={() => { this.setState({ counter: this.state.counter + 1 }); }}>Increment this counter
                </button>

            </div>
        );
    }
}
