import { parser } from './argv-parser.js'
import { echo, series, step } from './fizz-buzz.js'


const {nFrom, nTo, print} = parser()

// const start = Date.now()
const values = series(nFrom, nTo)
// const stop = Date.now()

if(print=='series') {
    echo(values)
} else {
    console.info(nTo, step(nTo))
}


// console.info(`${(stop - start)/1000} seconds`)
