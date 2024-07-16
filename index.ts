const a: boolean[][] = [[true, false, true, true], [false, true, false, true]];

class Game {
    constructor(private board: boolean[][]) {
    }

    public showGrid(): void {
        let grid: string = "";

        this.board.map((arr) => {
            arr.map((item) => {
                grid += item ? "x" : "o"
            });
            grid += "\n"
        });

        //@ts-ignore
        console.log(grid);
    }
}

const game = new Game(a);
game.showGrid();
