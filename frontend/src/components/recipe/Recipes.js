import React, { Component } from 'react';

import RecipeCard from '~/src/components/recipe/RecipeCard';

import styles from '~/src/components/recipe/Recipes.css';

export default class Recipes extends Component {
  render() {
    const { recipes } = this.props;
    const recipeCards = recipes.map((recipe) => {
      return (
        <div key={recipe.id}>
          <RecipeCard recipe={recipe} />
        </div>
      );
    });

    return (
      <div>
        <h3>
          Recipes
        </h3>
        <div className='container'>
          <div className="card-deck">
            {recipeCards}
          </div>
        </div>
      </div>
    );
  }
}
