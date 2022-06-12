const defaultParameters = {
    nFrom: 1,
    nTo: 100,
    print: 'series'
}

export const parser = () => {
    const results = defaultParameters

    const argv = process.argv.slice(2).map(item => {
        const [param, value] = item.split('=')

        return { param, value }
    })

    argv.forEach(item => {
        if (results?.[item.param]) {
            results[item.param] = item.value
        }
    })

    // sanitize

    return {
        nFrom: parseInt(results.nFrom),
        nTo: parseInt(results.nTo),
        print: ['series', 'end'].includes(results.print) ? results.print : 'series'
    }
}