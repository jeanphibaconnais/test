import React, { useState, useEffect } from 'react';

export function TestHook() {
    const [ count, setCount ] = useState(0);

    // Similar to componentDidMount and componentDidUpdate:
    useEffect(() => {
        document.title = `You clicked ${count} times`;
    });

    return (
        <div>
            <h3>Test hook function</h3>
            <span>You clicked {count} times</span>
            <button onClick={() => setCount(count + 1)}>Click me</button>
        </div>
    );
}
