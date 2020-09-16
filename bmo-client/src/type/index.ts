export class Card {
    term: string
    definition: string

    constructor(t: string, d: string) {
      this.term = t;
      this.definition = d;
    }
}

export type Cards = Card[]

export interface RootState {
    terms: Cards;
}
