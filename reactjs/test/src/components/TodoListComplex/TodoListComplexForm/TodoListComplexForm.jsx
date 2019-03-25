import React, { Component } from 'react';

class TodoListComplexForm extends Component {

    constructor(props) {
        super(props);
        this.state = {
            currentElement : ''
        }
    }

    addTodo = () => {
        this.props.addTodoParent(this.state.currentElement);
        this.setState({currentElement: ''});
    }
   
    handleChange = (event) => {
        this.setState({currentElement : event.target.value});
    }

    render () {
        return (
            <div>
                <p>Current value : {this.state.currentElement}</p>
                <input type="text" value={this.state.currentElement} onChange={this.handleChange} />
                <button type="submit" onClick={(event) => {
                     this.addTodo(event.target.value);
                }}>Add element</button>
            </div>
        );
    }

}


//     let currentState = '';

//     handleChange = (event) => {
//         this.setState({currentElement : event.target.value});
//     }

//     return <div>
//         <input type="text" value={currentState} onChange={this.handleChange} />
//         <button type="submit" onClick={(event) => {
//             addTodo(event.target.value);
//         }}>Add element</button>

//     </div>;
// }

export default TodoListComplexForm;