<div align="center">
    <img src="assets\spongebob.jpg" width="400" />
    <h2>SpongeBob Desktop</h2>
</div>

## Project Overview

SpongeBob Desktop is a lightweight background client that allows seamless communication between your desktop and mobile devices. It integrates functionalities such as controlling the desktop remotely, checking mobile connection status, and setting a secret code for secure interaction.

## Features

- **Background Operation**: The app runs silently in the background, occupying minimal system resources.
- **Mobile Connection Monitoring**: Allows checking if the mobile device is connected.
- **Secret Code**: Generates a random secret code for secure interaction, which can be reset at any time.
- **System Tray Icon**: Provides a system tray icon with options for managing mobile connections, secret codes, and more.

## Installation

To set up the SpongeBob Desktop client, follow these steps:

1. **Clone the repository**:

   ```bash
   git clone https://github.com/<your-username>/spongebob-server.git
   cd spongebob-server
   ```

2. **Build the project**:

   If you haven't already, you'll need Go installed on your system. You can download it from [here](https://golang.org/dl/).

   Once you have Go set up, you can build the project using the following command:

   ```bash
   go build
   ```

3. **Run the application**:

   After building the project, run the compiled executable:

   ```bash
   ./spongebob-server
   ```

4. **System Tray**:

   Upon running, SpongeBob Desktop will minimize to the system tray. You can interact with it through the tray icon.

## TODO

- [ ] Implement mobile-to-desktop communication
- [ ] Improve tray icon appearance and user feedback
- [ ] Add error handling for failed connections
- [ ] Implement better secret code reset mechanism
- [ ] Refactor the code to allow dynamic configuration loading

## Contributing

Contributions are welcome! If you have ideas or improvements, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
