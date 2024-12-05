const fs = require('fs');

function isCorrectUpdate (update, dictionary) {
    for (let j = 0; j < update.length; j++) {        
        let elem = update[j];

        if (elem in dictionary) {            
            for (let l = j-1; l >= 0; l--) {
                if (dictionary[elem].includes(update[l])) {
                    return false;
                }
            }
        }
     }

     return true;
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

    console.log(firstPart);

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

    // Output results
    // console.log('Dictionary:', dictionary);
    // console.log('2D Array:', array2D);

    let sum = 0;

    for (let i = 0; i < array2D.length; i++) {
        if (isCorrectUpdate(array2D[i], dictionary)) {            
            const midIdx = Math.ceil((array2D[i].length)/2) - 1
            sum = sum + array2D[i][midIdx];
        }
    }

    console.log(sum);
}

main();
