import { parser } from "./argv-parser"

test('no arguments gives default', ()=> {
    process.argv = ['dummy1', 'dummy2']

    const argv = parser()

    expect(argv).toStrictEqual({
        nFrom: 1,
        nTo: 100
    })
})

test('extra argument gives default', ()=> {
    process.argv = ['dummy1', 'dummy2', 'dummy=argument']

    const argv = parser()

    expect(argv).toStrictEqual({
        nFrom: 1,
        nTo: 100
    })
})

test('extra argument gives default', ()=> {
    process.argv = ['dummy1', 'dummy2', 'nFrom=-5']

    const argv = parser()

    expect(argv).toStrictEqual({
        nFrom: -5,
        nTo: 100
    })
})