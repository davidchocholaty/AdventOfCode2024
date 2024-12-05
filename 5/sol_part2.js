const fs = require('fs');

function isCorrectUpdate(update, dictionary) {
    for (let j = 0; j < update.length; j++) {
        let elem = update[j];
        if (elem in dictionary) {
            for (let l = j - 1; l >= 0; l--) {
                if (dictionary[elem].includes(update[l])) {
                    return false;
                }
            }
        }
    }
    return true;
}

function sortUpdate(update, dictionary) {
    return update.sort((a, b) => {
        if (a === b) return 0; // Handle equal elements
        if (dictionary[a] && dictionary[a].includes(b)) return -1;
        if (dictionary[b] && dictionary[b].includes(a)) return 1;
        return 0;
    });
}

function main() {
    if (process.argv.length !== 3) {
        console.log('Please provide exactly one file as a parameter.');
        return;
    }

    const filePath = process.argv[2];
    let fileContent = "";

    try {
        fileContent = fs.readFileSync(filePath, 'utf-8');
    } catch (error) {
        console.error(`Error reading file: ${error.message}`);
        return;
    }

    // Split the content into two parts by the empty line
    const [firstPart, secondPart] = fileContent.split('\n\n');

    // Parse the first part into a dictionary
    const dictionary = {};
    const firstLines = firstPart.trim().split('\n');
    for (const line of firstLines) {
        const [key, value] = line.split('|').map(Number); // Convert to numbers
        if (!dictionary[key]) {
            dictionary[key] = [];
        }
        dictionary[key].push(value);
    }

    // Parse the second part into a 2D array
    const array2D = secondPart
        .trim()
        .split('\n')
        .map(line => line.split(',').map(Number)); // Convert each value to a number

    let sumCorrect = 0;
    let sumIncorrect = 0;
    const incorrectUpdates = [];

    for (let i = 0; i < array2D.length; i++) {
        const update = array2D[i];
        if (isCorrectUpdate(update, dictionary)) {
            const midIdx = Math.floor(update.length / 2);
            sumCorrect += update[midIdx];
        } else {
            incorrectUpdates.push(update);
        }
    }

    for (const update of incorrectUpdates) {
        const sortedUpdate = sortUpdate(update, dictionary);
        const midIdx = Math.floor(sortedUpdate.length / 2);
        sumIncorrect += sortedUpdate[midIdx];
    }

    console.log("Part 1:", sumCorrect);
    console.log("Part 2:", sumIncorrect);
}

main();
