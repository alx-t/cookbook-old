import React, { Component } from 'react';

import { RECIPES } from '~/src/constants/Recipes';
import Recipes from '~src/components/recipe/Recipes';

export default class RecipesPage extends Component {
  render() {
    return (
      <Recipes recipes={RECIPES} />
    );
  }
}
