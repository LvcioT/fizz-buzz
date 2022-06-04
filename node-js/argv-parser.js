const defaultParameters = {
    nFrom: 1,
    nTo: 100
}

export const parser = () => {
    const results = defaultParameters

    const argv = process.argv.slice(2).map(item => {
        const [param, value] = item.split('=')

        return { param, value }
    })

    argv.forEach(item => {
        if(results?.[item.param]) {
            results[item.param] = parseInt(item.value)
        }
    })

    return results
}