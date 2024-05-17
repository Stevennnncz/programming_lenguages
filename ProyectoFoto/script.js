document.getElementById('uploadForm').addEventListener('submit', async (e) => {
    e.preventDefault();

    const fileInput = document.getElementById('imageInput');
    const file = fileInput.files[0];

    if (!file) {
        return; // No file selected, do nothing
    }
    
    const formData = new FormData();
    formData.append('image', file);
    
    try {
        const response = await fetch('/upload', {
            method: 'POST',
            body: formData
        });
        const data = await response.json();
        if (data.success) {
            document.getElementById('message').innerText = data.message;
            document.getElementById('message').classList.remove('hidden');

            // Show uploaded image
            const uploadedImage = document.getElementById('uploadedImage');
            uploadedImage.src = URL.createObjectURL(file);
            document.getElementById('imageContainer').classList.remove('hidden');
        } else {
            document.getElementById('message').innerText = data.message;
            document.getElementById('message').classList.remove('hidden');
            document.getElementById('imageInput').value = ''; // Clear the file input field
        }
    } catch (error) {
        console.error('Error:', error);
        document.getElementById('message').innerText = 'An error occurred while uploading the image.';
        document.getElementById('message').classList.remove('hidden');
    }
});
