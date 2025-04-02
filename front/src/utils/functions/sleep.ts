export const sleep = async (time = 2000): Promise<void> => {
    return new Promise((res) => {
        setTimeout(() => {res()}, time)
    })
}