package ar.com.ariel17.xy;

import android.content.Intent;
import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.view.View;
import android.widget.Button;
import android.widget.TextView;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        Button ok = (Button) findViewById(R.id.okButton);
        ok.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                TextView pin = (TextView) findViewById(R.id.pinText);

                Intent registrationIntent = new Intent(MainActivity.this, RegistrationActivity.class);
                registrationIntent.putExtra(RegistrationActivity.PIN_KEY, pin.getText().toString());

                MainActivity.this.startActivity(registrationIntent);
            }
        });

    }
}
