# Find the actual activation date of the phone number
### Step to run the app:
    git clone https://github.com/hungranger/phone-activation.git
    cd phone-activation/cmd/cli/
    go run main.go phone_numbers.csv output.csv

#### Note:
- If the input file is not provided ('phone_number.csv'). The app will use the default filepath in the config file
![](https://i.imgur.com/nSoPhrb.png)
- If the ouput file is not provided ('output.csv'). The app will generate will generate the output file with name 'output.csv' in the same directory with the application

