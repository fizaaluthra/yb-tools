#!/bin/bash

# Function to display error message and exit
function error_exit() {
  echo "$1"
  exit 1
}

echo "Enter the Core dump file name: "
read -r file_name

# Check if the Core file exists
if [ ! -f "$file_name" ]; then
  error_exit "Error: The file '$file_name' does not exist."
fi

# Check if the Core file is a valid core dump
file_type=$(file "$file_name")
if echo "$file_type" | grep -q "core file"; then
  echo "Great! The core dump file provided is a valid core dump, proceeding with analysis of core dump."
else
  error_exit "Error: The Core dump file is NOT an ELF core dump, please provide a valid core dump file. Make sure the file IS NOT compressed, if so please extract it and then try again!"
fi

# Extract the yugabyte binary version information from the core dump file
yb_db_version=$(strings "$file_name" | grep -E "/yugabyte/yb-software/yugabyte-" | head -n 1) || yb_db_version=""

yb_db_tar_file=$(echo "$yb_db_version" | awk -F "/" '{for (i=1; i<=NF; i++) if ($i == "yugabyte" && $(i+1) == "yb-software") print $(i+2)}' | sed 's/centos/linux/' | sed 's/$/.tar.gz/')

version=$(echo "$yb_db_tar_file" | awk -F "-" '{print $2}')

# Generate the URL for the yb-db binary
yb_db_tar_url="https://downloads.yugabyte.com/releases/$version/$yb_db_tar_file"
echo "The final URL is: $yb_db_tar_url"

# Download the yb db binary tar file

download_file="$yb_db_tar_file"
yb_db_install_dir="/home/yugabyte/yb-software"

if [ -f "$yb_db_install_dir/$yb_db_tar_file" ]; then
  echo "The file $yb_db_tar_file already exists in $yb_db_install_dir. Skipping the download step."
else
  echo "Downloading the YB version file to $yb_db_install_dir/$yb_db_tar_file"

  curl -L -# "$yb_db_tar_url" -o "$yb_db_install_dir/$yb_db_tar_file"
  if [ $? -eq 0 ]; then
    echo "Download of YB version file succeeded."
  else
    error_exit "Error: Download of YB version file failed."
  fi
fi

#Renamed executable dir, means we are going to extract the tar file binary in the same version format as core file.

executable_dir=$(echo "$yb_db_version" | awk -F "/" '{for (i=1; i<=NF; i++) if ($i == "yugabyte" && $(i+1) == "yb-software") print $(i+2)}')


# Function to display a rotating symbol
function spinner {
    local pid=$!
    local delay=0.75
    local spinstr='|/-\'
    while [ "$(ps a | awk '{print $1}' | grep $pid)" ]; do
        local temp=${spinstr#?}
        printf " [%c]  " "$spinstr"
        local spinstr=$temp${spinstr%"$temp"}
        sleep $delay
        printf "\b\b\b\b\b\b"
    done
    printf "    \b\b\b\b"
}

# Extract the binary
tar_file="$yb_db_install_dir/$yb_db_tar_file"
extracted_directory="$yb_db_install_dir/$executable_dir"

if [ -d "$extracted_directory" ]; then
  echo "$extracted_directory already exists, not extracting again."
else
  echo "Extracting $tar_file in $yb_db_install_dir"
  mkdir -p $yb_db_install_dir/$executable_dir && tar -xzf "$tar_file" --strip-components=1 -C $_ &>/dev/null &
  spinner
fi

# Execute post install script
post_install="$yb_db_install_dir/$executable_dir/bin/post_install.sh"

if [ -f "$post_install" ]; then
  echo "Executing post_install script to setup the binary as per core dump. Please bear with me!"
  $post_install &>/dev/null &
  spinner
else
  error_exit "Error: $post_install not found."
fi



# Replace the path in the yb_db_version with the new target directory
new_yb_db_version=$(echo "$yb_db_version" | awk -F "/yugabyte/yb-software" '{print "/yugabyte/yb-software"$2}' | sed "s|/yugabyte/yb-software/yugabyte-[^/]*|$yb_db_install_dir/$executable_dir|")
# Use the lldb command with the new input string
# Ask user to enetr available lldb command option for ease.

#The below section is to ask few more user inputs and redirection etc.



echo "Select an option for lldb command and press ENTER:"
echo "1. bt"
echo "2. thread backtrace all"
echo "3. Other lldb command"
echo "4. Quit"
read -r option

while [[ ! "$option" =~ ^(1|2|3|4)$ ]]; do
  echo "Error: Invalid option selected. Please select either 1, 2, 3 or 4."
  echo "Select an option for lldb command and press ENTER:"
  echo "1. bt"
  echo "2. thread backtrace all"
  echo "3. Other lldb command"
  echo "4. Quit"
  read -r option
done

if [ "$option" == "1" ]; then
  lldb_command="bt"
elif [ "$option" == "2" ]; then
  lldb_command="thread backtrace all"
elif [ "$option" == "3" ]; then
  echo "Enter the lldb command:"
  read -r lldb_command
fi

if [ "$option" != "4" ]; then
  echo "Do you want to redirect the output to a file? (y/n)"
  read -r redirect_output
  while [[ ! "$redirect_output" =~ ^(y|n)$ ]]; do
    echo "Error: Invalid option selected. Please enter either y or n."
    echo "Do you want to redirect the output to a file? (y/n)"
    read -r redirect_output
  done


if [ "$redirect_output" == "y" ]; then
  output_file="${file_name}_$(echo "$lldb_command" | tr -s ' ' '_')_analysis.out"
  echo "Output will be saved to $output_file"
  lldb -f "$new_yb_db_version" -c "$file_name" -o "$lldb_command" -o "quit"> "$output_file"
  echo "Analysis complete, the file '$output_file' has been saved."
else
  lldb -f "$new_yb_db_version" -c "$file_name" -o "$lldb_command"
fi
fi
echo "Exiting."
exit 0
