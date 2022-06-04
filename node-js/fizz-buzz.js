export const series = (from, to) => {
    const results = []

    for (let i = from; i <= to; i++) {
        results.push({
            step: i,
            value: step(i)
        })
    }

    return results
}

export const step = num => {
    if (num % 3 == 0 && num % 5 == 0) {
        return 'FizzBuzz'
    }

    if (num % 3 == 0) {
        return 'Fizz'
    }

    if (num % 5 == 0) {
        return 'Buzz'
    }

    return num
}

export const  echo  = values => {
    values.forEach(({step, value}) => console.info(step, value))
}