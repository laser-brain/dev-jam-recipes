export interface IRecipe {
  id: number
  name: string
  thumbnail: string
  description: string
  type: string
  difficulty: string
  servings: number
  ingredients: IIngredient[]
}

export interface IIngredient {
  name: string
  amount: string
}
