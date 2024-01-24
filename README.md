# Password Manager CLI

This is for the cli password manager "npsg", a cli tool used for making and managing passwords.

Made by Nikechukwu Okoronkwo

## Installation

You can install the **npsg** tool from the package archives. Check the link or check below for downloadable versions (tar.gz)

In order to set up the **npsg** tool, you will need to unzip and extract the tarball from your terminal.
```bash
tar -xzvf <file>.tar.gz
```

When you are done, export the tool to your "PATH" variable in order to make use of the command.
```bash
export PATH=$PATH:/bin # Temporary
echo 'export PATH=$PATH:/bin' >> <shell-config-file> # Permanent - Shell config file could be '~/.zshrc', '~/.bashrc' etc
```

For those of us making use of windows, you will need to manually unzip the file using your chosen software, and then 

## Usage
The **npsg** tool has some basic features for this release, but if you have any contributions as to what you would like to see next, you can do so in the appropriate channels, or email me at [this email](https://nikechukwu@gmail.com).
In order to get basic usage, just parse the following command
```bash
npsg --help
```

The basic functionalities are shown below:
```bash
npsg make # Generates password

npsg save <pswrd> # Saves the password <pswrd> with given reference to storage

npsg get # Gets password from storage

npsg config # Configure password storage (either only root uses it or not)
```

## Contributing

Contributing is welcome and will be looked over frequently. If you want to make any software based off of this tool, please ensure to give me appropriate credit.

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.


## License

[MIT](https://choosealicense.com/licenses/mit/)

## Contributors
- Nikechukwu Okoronkwo