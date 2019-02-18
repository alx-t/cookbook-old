import React from 'react';

// import styles from '~/src/components/recipe/RecipeCard.css';

export default class RecipeCard extends React.Component {
  render() {
    return (
      <div className="card" style={{'maxWidth': '20rem'}}>
        <img className="card-img-top" src="https://source.unsplash.com/featured/" alt="Card image top"/>
        <h3 className="card-title">Card title</h3>
        <h4 className="card-subtitle">Card subtitle</h4>
        <p className="card-text">This is a simple Card example</p>
      </div>
    );
  }
}
