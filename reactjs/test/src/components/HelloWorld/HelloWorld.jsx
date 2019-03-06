import React, { Component } from 'react';
import ListTerms from '../ListsTerm/ListTerm';

export default class HelloWorld extends Component {
    constructor(props) {
        super(props);
        this.state = {
            online: true,
            counter: 0,
            term: null,
            items: []
        };
    }

    callbackclick() {
        console.log('erreur ...');
    }

    onSubmit = (event) => {
        event.eventPreventDefault(); // permet de bloquer le fonctionnement normal du formulaire
        this.setState({
            term: '',
            items: [ ...this.state.items, this.state.term ]
        });
    };

    render() {
        const user = this.props.user;
        return (
            <div>
                <p>
                    {' '}
                    Hello {user}, you are {!this.state.online || 'online'} !{' '}
                </p>
                <p>{this.state.counter}</p>
                <button
                    onClick={() => {
                        this.setState({ counter: this.state.counter + 1 });
                    }}
                >
                    increment counter
                </button>

                <form onSubmit={this.onSubmit}>
                    <input value={this.state.term} onChange={this.onChange} />
                </form>

                <ListTerms items={this.props.items} />
            </div>
        );
    }
}
