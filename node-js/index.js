import { parser } from './argv-parser.js'
import { echo, series, step } from './fizz-buzz.js'


const {nFrom, nTo} = parser()

const start = Date.now()
const values = series(nFrom, nTo)
const stop = Date.now()

echo(values)

console.info(`${(stop - start)/1000} seconds`)

console.info(nTo, step(nTo))