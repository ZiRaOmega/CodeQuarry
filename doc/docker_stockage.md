When you use Docker volumes in a `docker-compose.yml` file, Docker manages the location of the volume on the host machine. To find out where the `img` folder's data is stored on your machine, you can inspect the Docker volumes.

Here's how you can determine the location:

1. **Identify the Volume Name**:
   The volume name is defined in the `docker-compose.yml` file. In this case, it's `img_data`.

2. **Inspect the Volume**:
   Use the `docker volume inspect` command to find the details about the volume, including its location on the host machine.

### Steps to Find the Volume Location

1. **List Docker Volumes**:
   Run the following command to list all Docker volumes:

   ```sh
   docker volume ls
   ```

   You should see an entry for `img_data`.

2. **Inspect the Volume**:
   Run the following command to inspect the `img_data` volume:

   ```sh
   docker volume inspect img_data
   ```

3. **Locate the Volume**:
   The output of the `inspect` command will include the `Mountpoint` field, which shows the location of the volume on the host machine. It will look something like this:

   ```json
   [
       {
           "CreatedAt": "2023-05-18T00:00:00Z",
           "Driver": "local",
           "Labels": {},
           "Mountpoint": "/var/lib/docker/volumes/img_data/_data",
           "Name": "img_data",
           "Options": {},
           "Scope": "local"
       }
   ]
   ```

   The `Mountpoint` field shows the path on the host machine where Docker stores the data for the `img_data` volume.

### Example Output

Assuming the `Mountpoint` is `/var/lib/docker/volumes/img_data/_data`, the `img` folder's data is stored in that directory on your host machine.

### Summary

To summarize, the `img` folder's data is stored in a Docker-managed volume. You can find the exact location on your host machine by using the `docker volume inspect` command and looking at the `Mountpoint` field in the output.

By following these steps, you can easily locate where Docker is storing the persistent data for the `img` folder on your machine.