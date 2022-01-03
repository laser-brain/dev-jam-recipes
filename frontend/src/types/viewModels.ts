export interface IRecipe {
  id: number;
  name: string;
  thumbnail: string;
  description: string;
  type: string;
  difficulty: string;
  ingredients: IIngredient[];
}

export interface IIngredient {
  name: string;
  amount: string;
}
