import React from 'react';

class HelloWorld extends React.Component {
    constructor(props) {
            super(props);
            this.state = {
                 online: false
            }
     }

     render() {
            const user = this.props.user;
            return (
                    <p> Hello {user}, you are {!this.state.online || 'online'} ! </p>
            )
      }
}