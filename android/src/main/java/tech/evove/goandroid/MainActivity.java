package tech.evove.goandroid;

import android.os.Bundle;
import android.support.design.widget.FloatingActionButton;
import android.support.v7.app.AppCompatActivity;
import android.support.v7.widget.Toolbar;
import android.util.Log;
import android.view.Menu;
import android.view.MenuItem;
import android.view.View;

import core.Counter;
import io.reactivex.Observable;
import io.reactivex.ObservableSource;
import io.reactivex.Scheduler;
import io.reactivex.android.schedulers.AndroidSchedulers;
import io.reactivex.functions.Consumer;
import io.reactivex.functions.Function;
import io.reactivex.schedulers.Schedulers;
import io.reactivex.subjects.PublishSubject;
import tech.evove.goandroid.core.GoScheduler;

import static java.lang.String.format;
import static java.util.Locale.US;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        Toolbar toolbar = (Toolbar) findViewById(R.id.toolbar);
        setSupportActionBar(toolbar);


        final PublishSubject<Boolean> subject = PublishSubject.create();
        FloatingActionButton fab = (FloatingActionButton) findViewById(R.id.fab);
        fab.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View view) {
                subject.onNext(true);
                subject.onNext(true);
                subject.onNext(true);
            }
        });

        setup(subject, false);
        setup(subject, true);
    }

    private void setup(PublishSubject<Boolean> subject, final boolean gogo) {
        final Counter counter = new Counter(0);
        Scheduler scheduler = gogo
                ? GoScheduler.instance()
                : Schedulers.io();
        subject.subscribeOn(scheduler)
                .observeOn(scheduler)
                .flatMap(new Function<Boolean, ObservableSource<Runnable>>() {
                    @Override
                    public ObservableSource<Runnable> apply(Boolean ignored) throws Exception {
                        final String txt = format(US, "Hello %d times!", counter.increment());
                        String prefix = gogo ? "gogo" : "nogo";
                        final String tag = prefix + "-" + Thread.currentThread().getName();
                        return Observable.<Runnable>just(new Runnable() {
                            @Override
                            public void run() {
                                Log.d(tag, txt);
                            }
                        });
                    }
                })
                .observeOn(AndroidSchedulers.mainThread())
                .subscribe(new Consumer<Runnable>() {
                    @Override
                    public void accept(Runnable r) throws Exception {
                        r.run();
                    }
                });
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu) {
        // Inflate the menu; this adds items to the action bar if it is present.
        getMenuInflater().inflate(R.menu.menu_main, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item) {
        // Handle action bar item clicks here. The action bar will
        // automatically handle clicks on the Home/Up button, so long
        // as you specify a parent activity in AndroidManifest.xml.
        int id = item.getItemId();

        //noinspection SimplifiableIfStatement
        if (id == R.id.action_settings) {
            return true;
        }

        return super.onOptionsItemSelected(item);
    }
}
