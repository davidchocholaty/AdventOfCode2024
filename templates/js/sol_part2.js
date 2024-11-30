const fs = require('fs');

function main() {
    if (process.argv.length !== 3) {
        console.log('Please provide exactly one file as a parameter.');
        return;
    }

    const filePath = process.argv[2];
    try {
        const fileContent = fs.readFileSync(filePath, 'utf-8');
        // console.log(fileContent);
    } catch (error) {
        console.error(`Error reading file: ${error.message}`);
    }
}

main();
